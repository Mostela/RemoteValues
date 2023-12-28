import os

import requests
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()


def startup_remote_config():
    return requests.get(os.getenv('REMOTE_CONFIG_URL')).json()


remote_config = startup_remote_config()


class DataRemote(BaseModel):
    key: str
    value: str


@app.get("/")
async def root():
    return {"person_name": remote_config['person']}


@app.get("/healthcheck")
async def healthcheck():
    return {"status": "ok"}


@app.post("/remoteconfig")
async def remoteconfig(data_remote: DataRemote):
    remote_config.__setitem__(data_remote.key, data_remote.value)
    return {"status": "ok"}
