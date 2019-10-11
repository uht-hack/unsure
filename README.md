Launching:        
 
The following 5 commands must be run, each in its own terminal window:

### Runs the engine: 
go run unsure/engine/engine/main.go --db_recreate --crash_ttl=0 --fate_p=0

### Runs each player (4 players in total):
go run uht-unsure/uht/main.go --db_recreate --uht_db="mysql://root@tcp(127.0.0.1:3306)/uht_1?" --engine_address="127.0.0.1:12048" --grpc_address="127.0.0.1:18000" --uht_cl1="127.0.0.1:18001" --uht_cl2="127.0.0.1:18002" --uht_cl3="127.0.0.1:18003" --player="1" --crash_ttl=0 --fate_p=0
        
go run uht-unsure/uht/main.go --db_recreate --uht_db="mysql://root@tcp(127.0.0.1:3306)/uht_2?" --engine_address="127.0.0.1:12048" --grpc_address="127.0.0.1:18001" --uht_cl1="127.0.0.1:18002" --uht_cl2="127.0.0.1:18003" --uht_cl3="127.0.0.1:18000" --player="2" --crash_ttl=0 --fate_p=0
        
go run uht-unsure/uht/main.go --db_recreate --uht_db="mysql://root@tcp(127.0.0.1:3306)/uht_3?" --engine_address="127.0.0.1:12048" --grpc_address="127.0.0.1:18002" --uht_cl1="127.0.0.1:18003" --uht_cl2="127.0.0.1:18000" --uht_cl3="127.0.0.1:18001" --player="3" --crash_ttl=0 --fate_p=0
         
go run uht-unsure/uht/main.go --db_recreate --uht_db="mysql://root@tcp(127.0.0.1:3306)/uht_4?" --engine_address="127.0.0.1:12048" --grpc_address="127.0.0.1:18003" --uht_cl1="127.0.0.1:18000" --uht_cl2="127.0.0.1:18001" --uht_cl3="127.0.0.1:18002" --player="4" --crash_ttl=0 --fate_p=0