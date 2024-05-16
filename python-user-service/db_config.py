import mysql.connector

mycon = mysql.connector.connect(
    host='127.0.0.1', 
    user='root',
    port="3306",
    password='fb123465',
    database='userdb', 
    auth_plugin="mysql_native_password")
# conn = mysql.connector.connect(user='root', password='fb123465', host='127.0.0.1',port=3306, database='userdb', auth_plugin='mysql_native_password')
print(mycon)
# Making MySQL cursor object
# mycur = mycon.cursor()
