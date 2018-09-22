import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import s from "../styles";
import UserList from "../components/UserList";
import {Permissions} from "expo";

export default class QuestionScreen extends Component {
  constructor(props) {
	super(props);
	this.state = {question: null, heading: null};
  }
	componentWillMount(){
		this._getLocationAsync();
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
			Expo.Location.watchHeadingAsync((obj) => {
				let heading = obj.magHeading;
				this.setState({ heading: heading })
			})
		}
	};
	componentWillUnmount() {
		Expo.Location.watchHeadingAsync();
	}

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
  render() {
	  //const { navigation } = this.props;
	  const place = this.state.question ? this.state.question.Place : null;
	  //const game_id = navigation.getParam('game_id', "N/A");
	  //const game_starting = this.state.starting_game;
	return (
	  <View style={s.container}>
		  {place?
		  ( <View>
			  <Text style={s.h1}>Missä on {JSON.stringify(place)}?</Text>
			  <Button
				  color='#7439A2'
				  title="Valmis"
                  onPress={() => {this.sendHeading()}}
			  />
		  </View>) : (
			  <Text>Ladataan...</Text>
			  )}



	  </View>
	);
  }
}

