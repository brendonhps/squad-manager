CREATE TABLE devs (
                      dev_id UUID PRIMARY KEY NOT NULL,
                      name VARCHAR(100) NOT NULL,
                      age INTEGER NOT NULL
);

CREATE TABLE requirement (
                             requirement_id UUID PRIMARY KEY NOT NULL,
                             dependencies TEXT []
);

CREATE TABLE deploy (
                        deploy_id UUID NOT NULL,
                        deploy_date date NOT NULL,
                        requirements_id UUID,
                        PRIMARY KEY (deploy_id,requirements_id)
);


CREATE TABLE project (
                         project_id UUID NOT NULL,
                         date_start date NOT NULL,
                         date_finish date NOT NULL,
                         date_homologation_start date NOT NULL,
                         date_homologation_end date NOT NULL,
                         devs TEXT [],
                         deploy_id UUID,
                         requirements_id UUID,
                         PRIMARY KEY (project_id,deploy_id,requirements_id)
);
