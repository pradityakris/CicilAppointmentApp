# Cicil Assigment Test
This is Assigment test creating Simple Appointment App Using Beego Framework.

# Instalation and Library used for this Apps to Work
- Please install these files on your current local GOPATH
1. [Beego (Beego Framework)](https://github.com/astaxie/beego)
2. [Bee (Beego Tools)](https://github.com/beego/bee)
3. [sqlite3 (Database used on this App)](https://github.com/mattn/go-sqlite3)

# Walkthrough
1. Make sure instalation and Libary installed on your current machine
2. Download this project
3. Extract Downloaded Project to "*cd yourgopath/src*"
4. Please rename extracted project to **CicilAppointmentApp**
4. Navigate to "**cd yourgopath/src/CicilAppointmentApp**"
5. Simply type this on your terminal "**bee run**" or "**go run main.go**"
6. At this stage the application should have been running on your local machine

# Feature
1. Get all meetup schedule (title, date, time) 
2. Get one meetup schedule (title, description, attendes,date, time,location)
3. create schedule
4. edit schedule
5. delete schedule
6. simple API for all the features above

# Steps
1. after Application successfully runs on your local machine
2. Test feature 1 *Get all meetup schedule (title, date, time)* on page http://localhost:8080/appointment/view
3. Test feature 2 and 4 *Get one meetup schedule (title, description, attendes,date, time,location)* by clicking **Edit** button on table
4. Test feature 3 *Create Schedule* by clicking **Add new Appointment** button and then fill the form and click **Create Appointment button**
5. Test feature 5 *Delete Schedule* by clicking **Delete** button

# API testing
- All the data retured are **JSON** type.
- You can test the api via Postman, dont forget to provide the params
1. Testing Get All meetup schedule **GET** API hit "http://localhost:8080/api/getallschedule"
2. Testing Get one meetup schedule **GET** API hit "http://localhost:8080/api/schedule/:id" please change **:id** to Appointment **Id**
3. Testing Create Schedule **POST** API hit "http://localhost:8080/api/addschedule?title=Testing API Create Schedule&date=07/11/2018&time=11:34 PM&desc=Testtest&atten=Rama Widhiantito&loc=Jakarta" the paramaters are **title, date, time, desc, atten, loc**
4. Testing Update Schedule **POST** API hit "http://localhost:8080/api/updateschedule/6?title=Testing API Create Schedule&date=07/11/2018&time=11:34 PM&desc=Testtest&atten=Rama Widhiantito&loc=Jakarta" please dont forget to add id on the url **ex : http://localhost:8080/api/updateschedule/6**
5. Testing Delete Schedule **GET** API hit "http://localhost:8080/api/deleteschedule/:id" please change **:id** to Appointment **Id**
