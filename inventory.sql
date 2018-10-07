create table inventory (
    id integer primary key autoincrement,
    name varchar(20) not null
);
create table hostgroup (
    id integer primary key autoincrement,
    name varchar(20) not null
);
create table host (
    id integer primary key autoincrement,
    name varchar(20) not null
);
create table inventory_hostgroups (
    inventory_id integer not null,
    hostgroup_id integer not null,
    foreign key (inventory_id) references inventory(id),
    foreign key (hostgroup_id) references hostgroup(id),
    primary key (inventory_id, hostgroup_id)
);
create table inventory_hosts (
    inventory_id integer not null,
    host_id integer not null,
    foreign key (inventory_id) references inventory(id),
    foreign key (host_id) references host(id),
    primary key (inventory_id, host_id)
);
create table hostgroup_hosts (
    hostgroup_id integer not null,
    host_id integer not null,
    foreign key (hostgroup_id) references hostgroup(id),
    foreign key (host_id) references host(id),
    primary key (hostgroup_id, host_id)
);
create table hostgroup_hostgroups (
    parent_hostgroup_id integer not null,
    child_hostgroup_id integer not null,
    foreign key (parent_hostgroup_id) references hostgroup(id),
    foreign key (child_hostgroup_id) references hostgroup(id),
    primary key (parent_hostgroup_id, child_hostgroup_id)
);
create table inventory_vars (
    inventory_id integer not null,
    var_name varchar(255) not null,
    var_value varchar(255) not null,
    foreign key (inventory_id) references inventory(id),
    primary key (inventory_id, var_name)
);
create table hostgroup_vars (
    hostgroup_id integer not null,
    var_name varchar(255) not null,
    var_value varchar(255) not null,
    foreign key (hostgroup_id) references hostgroup(id),
    primary key (hostgroup_id, var_name)
);
create table host_vars (
    host_id integer not null,
    var_name varchar(255) not null,
    var_value varchar(255) not null,
    foreign key (host_id) references host(id),
    primary key (host_id, var_name)
);