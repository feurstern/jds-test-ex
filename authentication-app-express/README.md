# authentication-app-express

To install dependencies: locally

```bash
bun install
```

To run:

```bash
bun run index.js
```

This project was created using `bun init` in bun v1.1.21. [Bun](https://bun.sh) is a fast all-in-one JavaScript runtime.



To run with docker:
I am asssume that you're currently using Linux, and my current machine is using arch linux

```bash
cd authentication-app-express
```

```bash
sudo docker compose up
```


![image](https://github.com/user-attachments/assets/3f4df0c3-b86c-4734-8976-4c9076a6052b)

// end if installation


API :

1. Create user
   Post ->  http://localhost:5454/auth/create

![image](https://github.com/user-attachments/assets/c3f359c8-822b-4b17-8da9-1ebab54af46e)


Copy paste the password from the terminal in order to test the next API
![image](https://github.com/user-attachments/assets/d35ddd95-1081-4eb7-99aa-c40d1d55ba51)



2. Login
   Post -> http://localhost:5454/auth/login
![image](https://github.com/user-attachments/assets/79951119-6994-4a01-ad84-d226bdf40a19)

3. Private-claim
   Post ->  http://localhost:5454/auth/private

   Ensure that you already copy paste the given token from the previous phase.
![image](https://github.com/user-attachments/assets/2d29495d-8e20-4901-88d8-0af84e978e12)

//end of api

   





