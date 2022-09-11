FROM python:3.10.6
WORKDIR /app
ADD . /app
RUN pip install --trusted-host pypi.python.org Flask
EXPOSE 9000
ENTRYPOINT ["python", "app.py"]
CMD ["--host=0.0.0.0"]

