import mysql.connector

mycon = mysql.connector.connect(
    host='dockerdbhost', 
    user='root',
    port="3306",
    password='fb123465',
    database='userdb')
# conn = mysql.connector.connect(user='root', password='fb123465', host='127.0.0.1',port=3306, database='userdb', auth_plugin='mysql_native_password')
print('mysql DB Connected')
# Making MySQL cursor object
# mycur = mycon.cursor()
