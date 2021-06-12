create table user_resource(

    uid int not null references user_info(uid) on update cascade on delete cascade ,
    rid int not null references resource_item(id) on update cascade on delete cascade,

    constraint urid primary key (uid, rid)
);
