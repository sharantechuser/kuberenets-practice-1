FROM python:3.9

WORKDIR /app

COPY ./dependencies.txt /app

RUN pip3 install -r dependencies.txt

COPY . .

EXPOSE 5000

CMD [ "python", "main.py" ]