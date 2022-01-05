# cara installasi
* Premis server
  - setup .env database koneksi(host, username, password, port, database name) pada case ini mengunakan postgresql
  - import sql script dari folder dbscript ke postgresql dengan nama file AmountBalance.sql
  - apabila library dari aplikasi belum terdownload dapat dilakukan dengan perintah "go mod download" 
  - build aplikasi golang dengan menjalankan perintah "go build -o api main.go"
  - jalankan aplikasi yang telah di build tadi dngan menggunakan perintah ./api
  - aplikasi juga dapat dijalankan tanpa harus mengkompile dengan perintah go run main.go

* Container
  - sesuaikan koneksi database dan aplikasi url yang akan dijalankan
  - jalankan dengan perintah "docker-compose up"
  - setelah postgres image telah terdownload dan berhasil dijalankan, langkah selanjutnnya bisa mengimport script sql yang ada di dalam folder dbscript dengan nama fileAmountBalance.sql

* Collection Postman
  - untuk collection apis bisa dilakukan dengan mengimport dari folder collections-postman dengan naama file "AccountBalance.postman_collection.json'

* Testing Aplikasi
  - untuk melakukan testing aplikasi bisa menjalankan perintah ./go_test.sh, apabila tidak tereksekusi maka dapat dilakukan dengan menambakan attibute execute,atau dapat dilakukan dengan menggukann perintah sh go_test.sh, di dalam go_testt.sh akan menjalankan test caase use case dan test case di repository.
