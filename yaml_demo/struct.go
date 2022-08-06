package yamldemo

/*
yaml 语法在线校验平台
https://www.bejson.com/validators/yaml_editor/
*/
var simple_cfg = `
database:
  dbtype: mysql
  dbname: narglc_database
  table: table
  username: username
  password: password
 
application:
  port: 8000
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type SimpleConf struct {
	Database    Database
	Application Application
}

type Database struct {
	Dbtype   string
	Dbname   string
	Table    string
	Username string
	Password string
}

type Application struct {
	Port string
}

/* 一个更复杂的列子 */
type Conf struct {
	Config Config
}

type Config struct {
	Models []Model
	Acls   []Acl
}

type Model struct {
	Name   string
	Schema string
}

type Acl struct {
	Model     string
	Role      string
	Operation string
	Action    string
}
