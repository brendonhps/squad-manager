CREATE TABLE devs (
                      dev_id INTEGER PRIMARY KEY NOT NULL,
                      name VARCHAR(100) NOT NULL,
                      age INTEGER NOT NULL
);

CREATE TABLE requirement (
                             requirement_id INTEGER PRIMARY KEY NOT NULL,
                             dependencies TEXT []
);

CREATE TABLE deploy (
                        deploy_id INTEGER NOT NULL,
                        deploy_date date NOT NULL,
                        requirements_id INTEGER,
                        PRIMARY KEY (deploy_id,requirements_id)
);


CREATE TABLE project (
                         project_id INTEGER NOT NULL,
                         date_start date NOT NULL,
                         date_finish date NOT NULL,
                         date_homologation_start date NOT NULL,
                         date_homologation_end date NOT NULL,
                         devs TEXT [],
                         deploy_id INTEGER,
                         requirements_id INTEGER,
                         PRIMARY KEY (project_id,deploy_id,requirements_id)
);
