// Function to check if the 'videos' database exists
function databaseExists(dbName) {
  var dbList = db.getMongo().getDBs();
  var databases = dbList.databases.map(function (database) {
    return database.name;
  });

  return databases.includes(dbName);
}

// Check if the 'videos' database already exists
if (!databaseExists("videos")) {
  // Create the 'videos' database
  db = db.getSiblingDB("videos");
  print('Creating the "videos" database...');
  // You can create collections, add initial data, or perform other operations here if needed
} else {
  print('The "videos" database already exists. Skipping creation.');
}
