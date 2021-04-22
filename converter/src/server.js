const express = require('express');
const multer = require('multer');
const { openBrowser, convertHtmlToPdf } = require('./puppeteerutil');

(async() => {
  const app = express();
  const upload = multer({
    dest: './tmp/uploads',
    limits: {
      fileSize: (Number(process.env.UPLOAD_FILESIZE_MB) || 5) * 1024 * 1024
    },
  });

  const browser = await openBrowser();

  try {
    app.post('/html-to-pdf', upload.single('html'), async (req, res) => {
      const html = req.file;
      console.log(html);
      const pdf = await convertHtmlToPdf(browser, html.path);

      res.writeHead(200, { 'content-type': 'application/pdf', });
      res.end(pdf, 'binary');
    });

    app.listen(8080, () => console.log('listening http://127.0.0.1:8080'));
  } catch (e) {
    await browser.close();
    console.error(e);
  }
})()
