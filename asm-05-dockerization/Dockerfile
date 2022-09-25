FROM python:3.10.6 AS build
WORKDIR /app
ADD . /app
RUN python -m venv /opt/venv
ENV PATH="/opt/venv/bint:$PATH"
RUN pip install --trusted-host pypi.python.org Flask



FROM python:3.10.6 AS applicationpy
WORKDIR /app
COPY --from=build /app/ .
ENV PATH=/root/.local/bin:$PATH

CMD [ "flask", "run", "--host=0.0.0.0", "--port=80" ]

