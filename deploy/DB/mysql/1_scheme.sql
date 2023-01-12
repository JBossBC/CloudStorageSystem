
create database if not exists cloudStorageSystem;
use cloudStorageSystem;

drop table  if exists FileMetaTable;
create table FileMetaTable(
    creator varchar(20) NOT NULL,
    createGroup varchar(20) NOT NULL,
    name varchar(30) NOT NULL,
    description varchar(30) NOT NULL ,    /*set the uri from this file or folder*/
    create_time timestamp default NOW(),
    authority varchar(3) NOT NULL DEFAULT 644,
    typeOf varchar(15) NOT NULL DEFAULT 'file',
    update_time timestamp NOT NULL DEFAULT NOW(),
    size bigint NOT NULL DEFAULT 0,
    isDir bool NOT NULL default  false,
    delete_time timestamp,
    primary key(creator,name),
    index(name,createGroup)
)ENGINE=myisam;

drop table if exists UserGroup;

create table UserGroup(
    groupName varchar(20),
    userName varchar(20),
    primary key(groupName,userName)
)engine=innodb;


