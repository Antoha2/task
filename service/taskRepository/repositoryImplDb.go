package repository

import (
	"errors"
	"fmt"
	"log"
	"strings"

	helper "github.com/antoha2/task"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
)

const countTask = 3 //максимальное кол-во записей для одного пользователя в базе

var (
	errNotFinedIdDB       = errors.New("id not fined")
	sqlConditionAllTasks  string
	sqlConditionUserTasks string
)

//Create
func (r *repositoryImplDB) Create(task *RepTask) error {

	var count int64
	if err := r.rep.Model(task).Where("user_id = ?", task.UserId).Count(&count).Error; err != nil {
		log.Println(err)
		return errNotFinedIdDB
	}

	if count < countTask {

		query := "INSERT INTO todolist (user_id, text, isdone) VALUES ($1, $2, $3) RETURNING task_id"
		result := r.rep.Table("todolist").Raw(query, task.UserId, task.Text, task.IsDone).Scan(&task.TaskId)
		if errors.Is(result.Error, gorm.ErrInvalidValue) {
			return errors.New("ошибка сознания задачи")
		}

		/* r.rep.Create(&task)
		log.Println(task.TaskId)
		if errors.Is(result.Error, gorm.ErrInvalidValue) {
			return errors.New("ошибка сознания задачи") */
		/* var id int
		query := "INSERT INTO todolist (user_id, text, isdone) VALUES ($1, $2, $3) RETURNING task_id"
		row := r.rep.Raw(query, task.UserId, task.Text, task.IsDone)
		if err := row.Scan(&id); err != nil {
			panic(err) */

		//task.TaskId = id
		log.Println("создана запись - ", task)
		return nil
	}

	errStr := fmt.Sprintf("не больше %d записей на пользователя", countTask)
	return errors.New(errStr)
}

//Read
func (r *repositoryImplDB) Read(readFilter *RepFilter) []RepTask {

	if readFilter.UserId != 0 {
		sqlConditionAllTasks = fmt.Sprintf(" user_id = %d", readFilter.UserId)
		sqlConditionUserTasks = fmt.Sprintf(" user_id = %d", readFilter.UserId)
	}

	sliceTask := make([]RepTask, 0)
	task := new(helper.Task)

	if readFilter.Ids == nil || len(readFilter.Ids) == 0 {

		stmtGet := r.rep.Where(sqlConditionAllTasks).Find(&sliceTask).Scan(&sliceTask)
		if stmtGet.RowsAffected == 0 {
			log.Println("записей нет")
		}
		log.Println("считано", len(sliceTask), "записей")
		log.Println(sliceTask)

		//msg := "записи считаны"

		r.logger.Write(sliceTask) //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

		return sliceTask
	}

	//strIds := strings.Trim(strings.Replace(fmt.Sprint(readFilter.Ids), " ", ",", -1), "[]")
	//strRead := fmt.Sprintf("SELECT * FROM todolist WHERE task_id IN (%s) %s", strIds, sqlConditionUserTasks)
	//stmtGet := r.rep.Raw(strRead).Scan(&sliceTask)

	strRead := fmt.Sprintf("task_id IN (%s)", strings.Trim(strings.Replace(fmt.Sprint(readFilter.Ids), " ", ",", -1), "[]"))
	stmtGet := r.rep.Where(strRead).Where(sqlConditionUserTasks).Find(&task).Scan(&sliceTask)
	if stmtGet.RowsAffected == 0 {
		log.Println("записей нет")
	}

	log.Println(sliceTask)
	return sliceTask
}

//Delete
func (r *repositoryImplDB) Delete(delTask *RepTask) error {

	stmtGet := r.rep.Where("task_id = ?", delTask.TaskId).First(&delTask)
	if stmtGet.RowsAffected == 0 {
		errStr := "запись с таким Id не найдена"
		log.Println(errStr)
		return errors.New(errStr)
	}

	if delTask.UserId != 0 {
		sqlConditionUserTasks = fmt.Sprintf(" AND user_id = %d", delTask.UserId)
	}
	strDelete := fmt.Sprintf("task_id = ? %s", sqlConditionUserTasks)
	result := r.rep.Where(strDelete, delTask.TaskId).Delete(&delTask)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("ошибка удаления задачи")
	}
	log.Printf("удалена запись с id - %v\n", delTask.TaskId)
	return nil
}

//Update
func (r *repositoryImplDB) Update(upTask *RepTask) error {

	if upTask.UserId != 0 {
		sqlConditionUserTasks = fmt.Sprintf(" AND user_id = %d", upTask.UserId)
	}
	strUpdate := fmt.Sprintf("task_id=? %s", sqlConditionUserTasks)
	result := r.rep.Table("todolist").Model(upTask).Where(strUpdate, upTask.TaskId).Select("text", "isdone").Updates(RepTask{Text: upTask.Text, IsDone: upTask.IsDone})
	if errors.Is(result.Error, gorm.ErrInvalidValue) {
		return errors.New("ошибка изменения задачи")
	}
	log.Println("изменена запись с id -", upTask.TaskId)
	return nil
}
