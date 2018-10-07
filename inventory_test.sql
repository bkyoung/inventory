insert into inventory values(null, "testInventory");
insert into inventory_vars values (1, "inventoryFoo", "inventoryBar");
insert into hostgroup values(null, "testHostGroup");
insert into hostgroup_vars values (1, "hgFoo", "hgBar");
insert into host values(null, "testHost");
insert into host_vars values (1, "foo", "bar");
insert into hostgroup_hosts values (1, 1);
insert into inventory_hosts values (1, 1);
insert into inventory_hostgroups values (1, 1);