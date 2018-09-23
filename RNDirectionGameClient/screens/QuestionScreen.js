import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import s from "../styles";
import UserList from "../components/UserList";
import {Permissions, Location} from "expo";

export default class QuestionScreen extends Component {
  constructor(props) {
	super(props);
	this.state = {question: null, heading: null, location: null};
  }
	componentWillMount(){
		this._getLocationAsync();
		this._getPositionAsync();
	}
	_getLocationAsync = async () => {
		// Checking device location permissions
		let { status } = await Permissions.askAsync(Permissions.LOCATION);
		if (status !== 'granted') {
			this.setState({
				errorMessage: 'Permission to access location was denied',
			});
		}
		else {
			Location.watchHeadingAsync((obj) => {
				let heading = obj.magHeading;
				this.setState({ heading: heading })
			})
		}
	};
	componentWillUnmount() {
		Location.watchHeadingAsync();
	}
    _getPositionAsync = async () => {
        let { status } = await Permissions.askAsync(Permissions.LOCATION);
        if (status !== 'granted') {
            this.setState({
                errorMessage: 'Permission to access location was denied',
            });
        } else {
        	console.log(status);
		}
		console.log("Fetching location");
        let location = await Location.getCurrentPositionAsync({});
        this.setState({ location: location });
        console.log(JSON.stringify(this.state.location));
    };
  componentDidMount() {
	const game_id = this.props.navigation.getParam('game_id', "N/A");
	this.getQuestion(game_id);
  }

  getQuestion(game_id) {
	fetch('https://techdays2018.appspot.com/api/games/'+game_id +'/questions/next', {
		method: 'GET',
		headers: {
			Accept: 'application/json',
			'Content-Type': 'application/json'
		}
	}).then((response) => response.json()).then((responseJson) => {
		this.setState({question: responseJson})
	});
  }
  sendResult(game_id, question_id, player_id, angle, location) {
      const latitude = location != null ? location.latitude: 0;
      const longitude = location != null ? location.longitude: 0;
      fetch('https://techdays2018.appspot.com/api/games/'+game_id+'/questions/'+question_id+'/answers', {
          method: 'POST',
          body: JSON.stringify({playerID: player_id, angle: angle, latitude: latitude, longitude: longitude}),
          headers: {
              Accept: 'application/json',
              'Content-Type': 'application/json'
          }
      }).then(() => {
          this.props.navigation.navigate("AnswerFeedbackScreen", {game_id: game_id, player_name: this.props.navigation.getParam('player_name', "N/A"), question_id: question_id});
      });
      console.log(game_id);
      console.log(question_id);
      console.log(player_id);
      console.log(angle);
      console.log(JSON.stringify(location));

  }
  render() {
	  //const { navigation } = this.props;
	  const place = this.state.question ? this.state.question.Place : null;
      const angle = this.state.heading;
      const location = this.state.location && this.state.location.coords ? this.state.location.coords : null;
      const question_id = this.state.question ? this.state.question.ID : null;
      const game_id = this.props.navigation.getParam('game_id', "N/A");
      const player_id = this.props.navigation.getParam('player_id', 0);

      //const game_id = navigation.getParam('game_id', "N/A");
	  //const game_starting = this.state.starting_game;
	return (
	  <View style={s.container}>
		  {place?
		  ( <View>
			  <Text style={s.h1}>Missä on {place.name}?</Text>
			  <Button
				  color='#7439A2'
				  title="Valmis"
                  onPress={() => {this.sendResult(game_id, question_id, player_id, angle, location)}}
			  />
		  </View>) : (
			  <Text>Ladataan...</Text>
			  )}



	  </View>
	);
  }
}

