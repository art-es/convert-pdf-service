const puppeteer = require('puppeteer');
const { readFileSync } = require('fs');

/**
 * Launch a new browser
 * @returns {Promise<puppeter.browser>}
 */
const openBrowser = async () => {
  return puppeteer.launch({
    executablePath: process.env.CHROME_BIN || null,
    browserName: 'chrome',
    args: ['--no-sandbox', '--headless', '--disable-gpu']
  });
}

/**
 * Generate a PDF file from HTML file
 * @param {puppeter.browser} browser
 * @param {string} htmlPath
 * @returns {Promise<Buffer>}
 */
const convertHtmlToPdf = async (browser, htmlPath) => {
  const buf = readFileSync(htmlPath, { encoding: 'utf-8' });
  const html = buf.toString('ascii');

  const page = await browser.newPage();
  await page.setContent(html, { waitUntil: 'networkidle2' });
  return await page.pdf({ format: 'A4', printBackground: true });
};

module.exports = { openBrowser, convertHtmlToPdf };
