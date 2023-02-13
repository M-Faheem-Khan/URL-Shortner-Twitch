import { useState } from "react";



const App = () => {
  const [originalURL, setOriginalURL] = useState("");
  const [shortURL, setShortURL] = useState("");

  const getShortURL = () => {
    const URL = "http://localhost:3000/api/v1/shorten"
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ "url": originalURL })
    }

    fetch(URL, requestOptions).then(response => response.json()).then((data) => {
      if (data.hasOwnProperty("error")) {
        console.error(data.error)
        setShortURL("Server error - try again?");
      } else {
        setShortURL("http://localhost:3000/" + data.shortID)
      }
    });
  }

  const ShowShortenURL = () => {

    if (shortURL !== "") {
      return <div className="mx-auto max-w-screen-sm text-sm text-left text-gray-500 newsletter-form-footer dark:text-gray-300">
      <strong>
        URL: 
      </strong>
      <a href={shortURL} className="font-medium text-primary-600 dark:text-primary-500 hover:underline">{shortURL}</a>.</div>
    }
  }


  return (
    <div className="grid h-screen place-items-center">
      <section className="bg-white dark:bg-gray-900">
        <div className="py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6 w-96">
          <div className="mx-auto max-w-screen-md sm:text-center">
            <h2 className="mb-4 text-3xl tracking-tight font-extrabold text-gray-900 sm:text-4xl dark:text-white">Shorten URL</h2>
            <p className="mx-auto mb-8 max-w-2xl font-light text-gray-500 md:mb-12 sm:text-xl dark:text-gray-400">Links live for 24 hours.</p>
            <form action="#">
              <div className="items-center mx-auto mb-3 space-y-4 max-w-screen-sm sm:flex sm:space-y-0">
                <div className="relative w-full">
                  <label htmlFor="email" className="hidden mb-2 text-sm font-medium text-gray-900 dark:text-gray-300">URL</label>
                  <input className="block p-3 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 sm:rounded-none sm:rounded-l-lg focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Enter url" type="url" id="email" required="" value={originalURL} onChange={e => setOriginalURL(e.target.value)} />
                </div>
                <div>
                  <button type="submit" className="py-3 px-5 w-full text-sm font-medium text-center text-white rounded-lg border cursor-pointer bg-primary-700 border-primary-600 sm:rounded-none sm:rounded-r-lg hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800" onClick={getShortURL} >Shorten</button>
                </div>
              </div>
              <ShowShortenURL />
            </form>
          </div>
        </div>
      </section>
    </div>
  );
}

export default App;
