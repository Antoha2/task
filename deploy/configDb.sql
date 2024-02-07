DROP TABLE IF EXISTS todolist;
create table todolist (
	Id INT not null,
	Text varchar(255) not null,
	IsDone bool not null default false	
)