import React, { Component } from 'react';
import TopMenu from './views/TopMenu/index';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      categories: [],
    };
  }

  // componentDidMount() {
  //   fetch('/api/v1/categories')
  //     .then(response => response.json())
  //     .then(categories => {
  //       console.log("CATEGORIES: ", categories.categories);
  //       return this.setState({categories: categories.categories})
  //     })
  //     .catch(error => console.log("FETCH ERROR: ", error));
  // }

  render() {
    const { categories } = this.state;
    return (
      <TopMenu />
    );
  }
}
export default App;
