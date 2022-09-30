FROM python:3.7-slim AS build
RUN apt-get update
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bint:$PATH"
COPY requirements.txt .
RUN pip install --user -r requirements.txt

FROM python:3.7-slim AS application
WORKDIR /flask
COPY --from=build /root/.local /root/.local
ENV PATH=/root/.local/bin:$PATH
COPY app.py .
CMD [ "flask", "run", "--host=0.0.0.0", "--port=80" ]