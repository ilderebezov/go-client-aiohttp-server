FROM library/python:3.9-slim

#RUN DEBIAN_FRONTEND=noninteractive apt-get update && apt-get install --no-install-recommends -y build-essential

RUN mkdir -p /aiohttp_files_hash
WORKDIR /aiohttp_files_hash

COPY . /aiohttp_files_hash

RUN python -m pip install --upgrade pip
RUN pip install poetry
RUN poetry config virtualenvs.create false && poetry install

#EXPOSE 8080

#CMD ["python", "-m", "src.app"]