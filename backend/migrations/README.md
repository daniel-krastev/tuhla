`cc=cockroach`

- start the node:               `cc start-single-node --insecure`

The command above would create the database in the current folder. This means that if you start the
execute the command again in different folder it would create a new database.

- connect:                      `cc sql --insecure`

- execute file:                 `cc sql <flags> --file <path_to_file>`
- execute statement:            `cc sql --execute="<sql statement>;<sql statement>" --execute="<sql-statement>" <flags>`  
