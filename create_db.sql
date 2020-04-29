CREATE DATABASE squadmanager;

CREATE TABLE project (
    project_id PRIMARY KEY NOT NULL,
    date_start date NOT NULL,
    date_finish date NOT NULL,
    date_homologation_start date NOT NULL,
    date_homologation_end date NOT NULL,
    devs TEXT [],
    FOREIGN KEY (deploy_id) REFERENCES deploy (deploy_id)
    FOREIGN KEY (requirements_id) REFERENCES requirements (requirements_id)
);

CREATE TABLE deploy (
    deploy_id PRIMARY KEY NOT NULL,
    deploy_date date NOT NULL,
    FOREIGN KEY (requirements_id) REFERENCES requirements (requirements_id)
);

CREATE TABLE requirement (
    requirement_id PRIMARY KEY NOT NULL,
    dependencies TEXT []
);

CREATE TABLE devs (
    dev_id PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    age INTEGER NOT NULL
);