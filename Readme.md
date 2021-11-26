# Shortlink using a GCP Cloud Run
# Serverless example

Make a Dir:

    mkdir shortlink && cd shortlink

Clone:

    git clone https://github.com/tonnytg/shortlink.git


Send to Cloud Run:

    gcloud run deploy


Send a request: <br/>
Change url request to yours:


    curl --location --request POST 'https://shortlink-jlattyrfza-uc.a.run.app' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "title":"test1",
            "body": "hello"
        }'
