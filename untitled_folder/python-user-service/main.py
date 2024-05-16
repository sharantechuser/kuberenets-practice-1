# import pymysql
from app import app
from db_config import mycon
from flask import jsonify, Blueprint
from flask import flash, request
# from werkzeug import generate_password_hash, check_password_hash
		
base_blueprint = Blueprint('base', __name__, url_prefix='/user')

@base_blueprint.route('/add', methods=['POST'])
def add_user():
	try:
		mycur = mycon.cursor()
		_json = request.json
		_name = _json['username']
		_email = _json['password']
		# _password = _json['pwd']
		# validate the received values
		if _name and _email and request.method == 'POST':
			#do not save password as a plain text
			# _hashed_password = generate_password_hash(_password)
			# save edits
			sql = "INSERT INTO user(username, password) VALUES(%s, %s)"
			data = (_name, _email,)
			# conn = mysql.connect()
			mycur.execute(sql, data)
			# conn.commit()
			mycon.commit()
			resp = jsonify('User added successfully!')
			resp.status_code = 200
			return resp
		else:
			return not_found()
	except Exception as e:
		print(e)
	finally:
		mycur.close()
		
@base_blueprint.route('/users')
def users():
	print('calling /users')
	try:
		mycur = mycon.cursor()
		# conn = mysql.connect()
		# mycur = conn.mycur(pymysql.cursors.DictCursor)
		mycur.execute("SELECT * FROM user")
		rows = mycur.fetchall()
		resp = jsonify(rows)
		resp.status_code = 200
		return resp
	except Exception as e:
		print(e)
	finally:
		mycur.close()


@base_blueprint.route('/user/<int:id>')
def user(id):
	try:
		mycur = mycon.cursor()
		# conn = mysql.connect()
		# mycur = conn.mycur(pymysql.cursors.DictCursor)
		mycur.execute("SELECT * FROM user WHERE user_id=%s", (id,))
		row = mycur.fetchone()
		resp = jsonify(row)
		resp.status_code = 200
		return resp
	except Exception as e:
		print(e)
	finally:
		mycur.close()
		

@base_blueprint.route('/userdetail/<string:username>/<string:password>')
def user_detail(username, password):
	try:
		mycur = mycon.cursor()
		# conn = mysql.connect()
		# mycur = conn.mycur(pymysql.cursors.DictCursor)
		print(username)
		mycur.execute("SELECT * FROM user WHERE username=%s and password=%s", (username,password,))
		row = mycur.fetchone()
		resp = jsonify(row)
		resp.status_code = 200
		return resp
	except Exception as e:
		print(e)
	finally:
		mycur.close()
		

@base_blueprint.route('/update', methods=['POST'])
def update_user():
	try:
		mycur = mycon.cursor()
		_json = request.json
		_id = _json['user_id']
		_name = _json['username']
		_password = _json['password']	
		# validate the received values
		print(_id)
		print(_name)
		print(_password)
		if _name and _password and _id and request.method == 'POST':
			#do not save password as a plain text
			# _hashed_password = generate_password_hash(_password)
			# save edits
			sql = "UPDATE user SET username=%s, password=%s WHERE user_id=%s"
			data = (_name, _password, _id,)
			# conn = mysql.connect()
			# mycur = conn.mycur()
			mycur.execute(sql, data)
			# conn.commit()
			mycon.commit()
			resp = jsonify('User updated successfully!')
			resp.status_code = 200
			return resp
		else:
			return not_found()
	except Exception as e:
		print(e)
	finally:
		mycur.close()
		
@base_blueprint.route('/delete/<int:id>')
def delete_user(id):
	try:
		mycur = mycon.cursor()
		# conn = mysql.connect()
		# mycur = conn.mycur()
		mycur.execute("DELETE FROM user WHERE user_id=%s", (id,))
		# conn.commit()
		mycon.commit()
		resp = jsonify('User deleted successfully!')
		resp.status_code = 200
		return resp
	except Exception as e:
		print(e)
	finally:
		mycur.close()
		
@base_blueprint.errorhandler(404)
def not_found(error=None):
    message = {
        'status': 404,
        'message': 'Not Found: ' + request.url,
    }
    resp = jsonify(message)
    resp.status_code = 404
    return resp
		

# Register the blueprint with the Flask application
app.register_blueprint(base_blueprint)

if __name__ == "__main__":
    app.run(port=8000)