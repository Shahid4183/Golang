### Create config.yaml file and add following configurations to it

- env - specifies environment configuration to specify which environment on we are running
- it also specifies whether to enable ssl mode or not
- if env = local then frontend host = localhost and ssl mode is disabled
- if env = test then frontend host = test server and ssl mode is enable
- if env = prod then frontend host = production server ssl mode is enable

```sh
env: specify-env
```

- db - database configurations
- host - specifies where database is hosted
- port - specifies on which port database is running
- name - database name
- user - username of user avaible in database
- password - password of user specified
- sslmode - specifies wether to user ssl to connect with database
- sslcertificate - specifies specify ssl rool certificate if ssl mode is other than disable

```sh
db:
  host: specify-database-host-name
  port: specify-database-host-port
  name: specify-database-name
  user: specify-database-username
  password: specify-database-password
  sslmode: specify-database-sslmode
```
- if sslmode is other than disable then specify ssl root certificate

```sh
  sslcertificate: specify-path-to-ssl-certificate 
```

- server - server configurations
- port - specifies port on which server is running
- ssl encryption will be enabled if environment configurations (env variable) is set to other than local
- sslkey - specifies path to ssl private key used access ssl certificate
- sslcertificate - specifies path to ssl certificate used to encrypt traffic to the server

```sh
server:
  port: specify-port
  sslkey: specify-root-to-ssl-key
  sslcertificate: specify-root-to-ssl-certificate
```  

- logger - logger file configurations
- filename - logger filename
- directory - path where logger file will be created and stored

```sh
logger:
  filename: specify-log-filename
  directory: path-to-logger-directory
```

- workers - worker pool configuration
- amount - specifies amount of workers created to handle concurrent tasks
- channelCapacity - specifies channel capacity

```sh
workers:
  amount: specifies-number-of-workers-to-create
  channelCapacity: specifies-channel-capacity-in-numbers
```
- other configurations can be added here