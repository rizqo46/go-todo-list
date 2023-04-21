package db

import "github.com/astaxie/beego/orm"

func Migrate(o orm.Ormer) {
	err := o.Begin()
	if err != nil {
		panic(err)
	}

	defer o.Rollback()

	_, err = o.Raw(`CREATE table IF NOT EXISTS activities (
		activity_id INT NOT NULL AUTO_INCREMENT,
		title varchar(255),
		email varchar(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (activity_id)
	);`).Exec()
	if err != nil {
		panic(err)
	}
	_, err = o.Raw(`CREATE table IF NOT EXISTS todos (
		todo_id INT NOT NULL AUTO_INCREMENT,
		activity_group_id INT NOT NULL,
		title varchar(255),
		priority varchar(255),
		is_active BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		PRIMARY KEY (todo_id),
		FOREIGN KEY (activity_group_id) REFERENCES activities(activity_id)
	);`).Exec()
	if err != nil {
		panic(err)
	}

	err = o.Commit()
	if err != nil {
		panic(err)
	}
}
