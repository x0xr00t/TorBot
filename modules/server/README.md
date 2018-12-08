TorBot Server

This is a WebSocket server that communicates with the front-end by executing TorBot functions and sending the data asynchronously to the front-end to be displayed. The reason we're using async is for instant response time rather than waiting for a costly function to finish, we just consume the results as we produce them in the backend. 

To start the server, you can either use `run_server.sh` inside of the `dev` folder or just run `python3 torBot.py --server` directly. You must start the server before being able to use the front-end
