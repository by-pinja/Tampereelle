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
        this.timer = setInterval(()=> this.getResults(game_id, question_id), 1000);
    }

    componentWillUnmount() {
        this.timer = null;
    }
  getResults(game_id, question_id ) {
    fetch('https://techdays2018.appspot.com/api/games/'+game_id +'/questions/'+question_id, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((response) => response.json()).then((responseJson) => {
        if(responseJson.answer){
            this.timer = null;
            this.setState({results: responseJson});
        }
    });
  }
  render() {
      const results = this.state.results;
      const players = this.results.players;
    return (
      <View style={s.container}>
        <Text style={s.h1}>Vastaukset</Text>
          {results ? (
              <Text>Odotetaan vastauksia...</Text>
          ) : (
              <UserResponsesList players={players}/>
          )}

      </View>
    );
  }
}

