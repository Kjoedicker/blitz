const 
    express = require('express'),
    app = express(),
    port = 3000
;

app.get('/', (req, res) => {
    console.count("For the sake of vision");
    return res.send(JSON.stringify(req.headers));
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
})