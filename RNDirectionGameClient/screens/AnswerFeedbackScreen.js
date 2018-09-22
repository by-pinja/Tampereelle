import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
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
    /*renderList(results){
        {return results.map(result => (
            <View key={result.player_name}>
                <Text>{result.player_name}</Text>
            </View>
        );}
    }*/
  render() {
      const { navigation } = this.props;
      const player_name = navigation.getParam('player_name', "N/A");
      const game_id = navigation.getParam('game_id', "N/A");
      const question_id = navigation.getParam('question_id', "N/A");
      const results = this.state.results;
    return (
      <View style={s.container}>
        <Text style={s.h1}>Vastaukset</Text>
          {results ? (
              <Text>Odotetaan vastauksia...</Text>
          ) : (
              <View>
                  {renderList(results)}
              </View>
              <Text>Vastaukset tähän</Text>
          )}

      </View>
    );
  }
}

