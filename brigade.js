const { events, Job } = require("brigadier")

events.on("push", function(e, project) {
    console.log("received push for commit " + e.commit)
  })