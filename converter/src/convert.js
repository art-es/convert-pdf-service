const { openBrowser, convertHtmlToPdf } = require('./puppeteerutil');

(async () => {
  const args = process.argv;
  if (args.length < 4) {
    console.error("html and pdf paths are required");
    process.exit(1);
  }

  const htmlPath = args[2];
  const pdfPath = args[3];

  const browser = await openBrowser();
  await convertHtmlToPdf(browser, htmlPath, pdfPath);
  await browser.close();
})();
