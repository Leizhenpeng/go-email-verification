db.createUser(
    {
        user: "admin",
        pwd: "forkway",
        roles: [
            {
                role: "readWrite",
                db: "email-verifications"
            }
        ]
    }
);
