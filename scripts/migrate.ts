import { migrate } from "drizzle-orm/postgres-js/migrator";
import { db } from "../packages/db/src/db";

async function main() {
    console.log("Running migrations...");
    try {
        await migrate(db, { migrationsFolder: "packages/db/drizzle" });
        console.log("Migrations completed successfully!");
        process.exit(0);
    } catch (err) {
        console.error("Error running migrations:", err);
        process.exit(1);
    }
}

main();
