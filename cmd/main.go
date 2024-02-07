package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	//"github.com/golang-migrate/migrate/v4/database/postgres"

	pgx "github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/antoha2/task/config"
	authService "github.com/antoha2/task/pkg/auth"
	logger "github.com/antoha2/task/pkg/logger"
	taskService "github.com/antoha2/task/service"
	taskRepository "github.com/antoha2/task/service/taskRepository"
	web "github.com/antoha2/task/transport"
)

func main() {

	Run()

}

func initDb(cfg *config.Config) (*gorm.DB, error) {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
		cfg.DB.Sslmode,
	)

	// Prep config
	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf(" failed to parse config: %v", err)
	}

	// Make connections
	dbx, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(connConfig))
	if err != nil {
		return nil, fmt.Errorf(" failed to create connection db: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbx,
	}), &gorm.Config{})

	err = dbx.Ping()
	if err != nil {
		return nil, fmt.Errorf(" error to ping connection pool: %v", err)
	}
	log.Printf("(task) Подключение к базе данных на http://127.0.0.1:%d\n", cfg.DB.Port)
	return gormDB, nil
}

func Run() {

	cfg := config.GetConfig()
	gormDB, err := initDb(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//addrAuthStr := fmt.Sprintf("http://%s:%d", cfg.Auth.Host, cfg.Auth.Port) //:8180

	logger := logger.NewLogger(cfg)
	if err = logger.Init(); err != nil {
		log.Fatalf("ошибка log.init: %s\n", err)
	}

	GRPCConn, err := web.StartGRPC()
	if err != nil {
		log.Fatalf("ошибка GRPCStart: %s\n", err)
		//log.Println(err)
	}

	taskRep := taskRepository.NewTaskRepository(gormDB, logger)

	auth := authService.NewAuthService(GRPCConn)
	task := taskService.NewTaskService(taskRep)

	Tran := web.NewWeb(task, auth)

	go Tran.StartHTTP()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit
	Tran.Stop()

}

/////////////////////////////////////////////////////////////////////////////

/* type ParseToken struct {
	Token string `json:"token"`
}

type GetRoles struct {
	Roles []string `json:"roles"`
} */

//var conn *grpc.ClientConn

/*
func GetRolesResp(Id int) ([]string, error) {

	client := pb.NewTaskServiceClient(config.GRPCConn)
	request := &pb.GetRolesRequest{
		Id: int32(Id),
	}
	response, err := client.GetRoles(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("2fail to dial: %v", err)
		return nil, err
	}

	return response.Roles, fmt.Errorf(response.Err)
}

func ParseTokenResp(Token string) (int, error) {

	client := pb.NewTaskServiceClient(config.GRPCConn)
	request := &pb.ParseTokenRequest{
		Token: Token,
	}
	response, err := client.ParseToken(context.Background(), request)

	if err != nil {
		grpclog.Fatalf("2fail to dial: %v", err)
		return 0, err
	}

	return int(response.Id), fmt.Errorf(response.Err)
}
*/
/* func HttpStart() error {
	//var err error
	server := &http.Server{Addr: ":8180"}

	mux := http.NewServeMux()
	mux.HandleFunc("/sum", handlerSum)
	log.Printf("Запуск веб-сервера на http://127.0.0.1:%s\n", server.Addr) //:8080
	if err := http.ListenAndServe(server.Addr, mux); err != nil {
		return err
	}
	return nil
}
*/
/*
func Decoder(r *http.Request, msg *msg) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, msg)
	if err != nil {
		fmt.Println("can't unmarshal: ", err.Error())
		return err
	}
	return nil
}
*/
/*
func handlerParseToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	token := ParseToken{}
	err := Decoder(r, &token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := ParseTokenResp(token.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(fmt.Sprintf("%d", id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)
}

func handlerGetRoles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	roles := GetRoles{}
	err := Decoder(r, &roles)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	roles, err := GetRolesResp(roles.Roles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json, err := json.Marshal(fmt.Sprintf("%d", id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(json)
}
*/

/*
if err := GRPCStart(); err != nil {
	log.Println(err)
}

if err := HttpStart(); err != nil {
	log.Println(err)
}
*/
