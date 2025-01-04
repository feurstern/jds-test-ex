const request = require("supertest");
const app = require("../index"); 
const bcrypt = require("bcrypt");

let generatePassword;
describe("POST /auth/create", () => {
  it("should create a user and return NIK, Role, and password", async () => {
    const response = await request(app)
      .post("/auth/create")
      .send({ NIK: "1234567890123456", Role: "Admin" });

    expect(response.statusCode).toBe(200);
    expect(response.body).toHaveProperty("NIK");
    expect(response.body).toHaveProperty("Role");
    expect(response.body).toHaveProperty("password");

    generatePassword = response.body.password;
    // console.log("password generated:", generatePassword);
  });
});

describe("POST /auth/login", () => {
  it("should login and return a token", async () => {
    const response = await request(app).post("/auth/login").send({
      NIK: "1234567890123456",
      password: generatePassword,
    });

    expect(response.statusCode).toBe(200);
    expect(response.body).toHaveProperty("token");
  });
});
