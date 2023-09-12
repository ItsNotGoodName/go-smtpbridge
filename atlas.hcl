// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://migrations/schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = "sqlite://smtpbridge_data/smtpbridge.db?_fk=1"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "sqlite://file?mode=memory&_fk=1"

  migration {
    // URL where the migration directory resides.
    dir = "file://migrations"
    // An optional format of the migration directory:
    // atlas (default) | flyway | liquibase | goose | golang-migrate | dbmate
    format = goose
  }
}
