import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import s from "../styles";
import UserList from "../components/UserList";

export default class QuestionScreen extends Component {
  constructor(props) {
    super(props);
    this.state = {game: null, starting_game: false};
  }
    componentDidMount() {
        const game_id = this.props.navigation.getParam('game_id', "N/A");
        this.timer = setInterval(()=> this.getGame(game_id), 2000);
    }

    componentWillUnmount() {
        this.timer = null;
    }
  getGame(game_id) {
    fetch('https://techdays2018.appspot.com/api/games/'+game_id, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        }
    }).then((response) => response.json()).then((responseJson) => {
        if(responseJson.state && responseJson.state === "STARTED"){
            this.props.navigation.navigate("QuestionScreen", {game_id: game_id, player_name: this.props.navigation.getParam('player_name', "N/A")});
        }
        else{
            this.setState({game: responseJson})
        }
    });
  }
  startGame(game_id){
      this.setState({starting_game: true});
      fetch('https://techdays2018.appspot.com/api/games/'+game_id +'/state', {
          method: 'PUT',
          body: JSON.stringify({
              state: 'STARTED',
          }),
          headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json'
          }
      });
  }
  render() {
      const { navigation } = this.props;
      const players = this.state.game && this.state.game.players ? this.state.game.players : [];
      const game_id = navigation.getParam('game_id', "N/A");
      const game_starting = this.state.starting_game;
    return (
      <View style={s.container}>
          <Text style={s.h1}>Peli: {game_id}</Text>
          <UserList players={players}/>
          <Button style={s.button}
              disabled={game_starting}
              onPress={() => {
                  this.startGame(game_id);
              }}
              color='#7439A2'
              title="Aloita peli"
          />
          <Text style={s.h2}>Odotetaan pelaajia</Text>
      </View>
    );
  }
}

