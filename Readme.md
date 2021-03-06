# Studenth
## _A website for students, by students_

Studenth is a website for french students, developped on the occasion of national hackaton
of My Digital School.

:star2: You can visit the website by clicking :point_right: [here](http://ec2-15-236-205-10.eu-west-3.compute.amazonaws.com) :point_left:

- :wave: Studenth is here to help :wave:
- :imp: Studenth is an heresy :imp:
- :blush: Studenth is not beautifull but nice :kiss:

## Features

- Search informations all around the website
- Search a student house
- Display all availlable student houses in France
- Search an information about social support
- Display many articles about social support and how to subscribe to it

## Tech

Studenth uses theses languages

- [Golang] - C for dummies
- [Gin] - A high-performance HTTP web framework written in Golang
- [Javascript] - To make the website dynamic again!
- [HTML5] - For the html templates
- [CSS3] - For beauty

## Installation

Studenth requires [MySQL](https://www.mysql.com) or [MariaDB](https://mariadb.org) to run.

- To install [MySQL](https://www.mysql.com), you can follow [this tutorial](https://dev.mysql.com/doc/mysql-installation-excerpt/5.7/en/)
- But if you prefer [MariaDB](https://mariadb.org) you can follow [this one](https://mariadb.com/kb/en/getting-installing-and-upgrading-mariadb/)

#### Optional
If you want to rebuild the exec file, you have to install [Golang](https://golang.org/) v1.12+ and
- To install [Golang](https://golang.org/), you can follow [this tutorial](https://golang.org/doc/install)

Firstly, you have to clone this repository:
```sh
cd /your-dir
git clone https://github.com/TavernierAlicia/MDS-hackathon-2021.git
```

### The database

Before importing the database, you have to create one with a similar name in your own service
```sh
mysql -u username –-password=your_password 
CREATE DATABASE studenth;
```

Now, import the file "database.sql" on your database service
```sh
cd /path/to/sql/file
mysql -u username –-password=your_password database_name < database.sql 
```

Then, start the service
```sh
sudo service mysql start
```

##### :warning: Warning: 
My config is maybe not yours, so you can configure the port and host of the database in "conf.json" file

### The program
Go to the folder "MDS-hackathon-2021"
```sh
cd MDS-hackathon-2021
```

Now you can launch the program compiled for Linux OS...
```sh
./main
```

... Or compile it for your OS

```sh
go build main.go router.go routes.go datastructs.go
./main
```
##### :warning: Warning: 
The app is serving on 127.0.0.1:3000, if you want to change the DNS, you have to do it manually in router.go at line 55

Now, open your web browser and go to http://127.0.0.1:3333/ and TADAAA! :fireworks: :fireworks: :fireworks:
