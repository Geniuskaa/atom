
CREATE TABLE car_info (
    id bigserial,
    plate varchar(9) unique,
    model varchar not null,
    year_of_manufacture date not null,
    type varchar CHECK ( type IN ('CARSHARING', 'TAXI', 'DELIVERY', 'UNDEFINED')),
    is_deleted bool default false,
    PRIMARY KEY (id, plate)
);


CREATE TABLE car_mileage (
    id bigserial,
    update_time timestamp default now(),
    plate varchar(9) references car_info(plate),
    mileage int not null check ( mileage > 0 )
);