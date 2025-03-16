package main
import "fmt"

type Database interface {
	Query() string
}

type MySQL struct {
	UserName string
	Password string
	id       int
}

func (m *MySQL) Query(query string) string {
	return fmt.Sprintf("%s Имя: %s Пароль: %s Айди: %d", query, m.UserName, m.Password, m.id)
}

type PostgreSQL struct {
	UserName string
	Password string
	id       int
}

func (p *PostgreSQL) Query(query string) string {
	return fmt.Sprintf("%s Имя: %s Пароль: %s Айди: %d", query, p.UserName, p.Password, p.id)
}


func main {
	
}