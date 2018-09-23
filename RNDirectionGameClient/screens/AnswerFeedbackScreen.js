import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import UserResponsesList from '../components/UserResponsesList';
import s from "../styles";

export default class AnswerFeedbackScreen extends Component {
  constructor(props) {
    super(props);
    this.state = {results: null};
  }
    componentDidMount() {
        const question_id = this.props.navigation.getParam('question_id', "N/A");
        const game_id = this.props.navigation.getParam('game_id', "N/A");
        const player_name = this.props.navigation.getParam('player_name', "N/A");
        this.timer = setInterval(()=> {this.getResults(game_id, question_id); console.log("polling2") }, 5000);
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }
  getResults(game_id, question_id ) {
    fetch('https://techdays2018.appspot.com/api/games/'+game_id +'/questions/'+question_id, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((response) => response.json()).then((responseJson) => {
        if(responseJson != null && responseJson.length){
            clearInterval(this.timer);
            this.setState({results: responseJson});
        }
    });
  }
  render() {
      const results = this.state.results;
    return (
      <View style={s.container}>
        <Text style={s.h1}>Vastaukset</Text>
          {results ? (
              <View>
                  <Text style={s.h2}>Pelaajat:</Text>
                  <UserResponsesList players={results}/>
              </View>
          ) : (
              <Text>Odotetaan vastauksia...</Text>
          )}

      </View>
    );
  }
}

