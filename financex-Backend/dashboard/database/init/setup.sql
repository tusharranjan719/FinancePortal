drop table if exists participant_expense;
drop table if exists expense;
drop table if exists participant;
drop table if exists billSplit;


create table billSplit (
                       id         serial primary key,
                       uuid       varchar(64) not null unique,
                       name       varchar(255),
                       created_at timestamp not null
);

create table participant (
                             id         serial primary key,
                             uuid       varchar(64) not null unique,
                             name      varchar(255) not null,
                             billSplit_id    integer references billSplit(id),
                             created_at timestamp not null,
                             CONSTRAINT U_Participant UNIQUE (name, billSplit_id)
);

create table expense (
                        id         serial primary key,
                        uuid       varchar(64) not null unique,
                        name       varchar(255) not null,
                        amount     float8 not null,
                        billSplit_id    integer references billSplit(id),
                        participant_id    integer references participant(id),
                        created_at timestamp not null
);

create table participant_expense (
                        id         serial primary key,
                        participant_id    integer references participant(id),
                        expense_id    integer references expense(id)
);


