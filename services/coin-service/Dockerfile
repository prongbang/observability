FROM python:3.10-slim-buster
WORKDIR /app
RUN python3 -m pip install grpcio-tools
COPY . .
EXPOSE 50053
CMD ["python3", "main.py"]