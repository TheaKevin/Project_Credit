import './App.css';
import {useState, useEffect} from 'react'

function App() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch("http://localhost:8080");
      const data = await res.json();
      setData(data);
    }
    catch (error) {
      setError(error);
    }
  }

  if (error) {
    return <div>Failed to load {error.toString()}</div>
  }

  if (!data) {
    return <div>Loading...</div>
  }

  return (
    <div>
      <p>data: {data.message}</p>
    </div>
  )
}

export default App;
