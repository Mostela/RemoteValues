FROM python
WORKDIR app/
LABEL authors="joao"
EXPOSE 8000
COPY requirements.txt app/
RUN pip install -r requirements.txt
COPY . app/
ENTRYPOINT uvicorn main:app --host 0.0.0.0 --port 8000
