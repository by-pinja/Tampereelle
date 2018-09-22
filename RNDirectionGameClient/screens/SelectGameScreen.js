import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert} from 'react-native';
import s from "../styles";

export default class SelectGameScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            game_id: ''
        };
    }
    createGame() {
        fetch('http://192.168.43.60:8080/api/games', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            }
        }).then((response) => response.json()).then((responseJson) => {
            Alert.alert("Message", JSON.stringify(responseJson))
        });
    }
    showGames() {
        fetch('http://192.168.43.60:8080/api/games/0', {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            }
        }).then((response) => response.json()).then((responseJson) => {
            Alert.alert("Message", JSON.stringify(responseJson))
        });
    }
    render() {
        const { navigation } = this.props;
        const player_name = navigation.getParam('player_name', "N/A");
        return (
            <View style={{padding: 10, display: 'flex', flex: 1}}>
                <Text style={s.h1}>Liity peliin </Text>
                <Text style={s.h2}>Pelaajan nimi: { player_name } </Text>
                <View style={{display: 'flex', flexDirection: 'row', paddingBottom: 10}}>
                    <TextInput
                        style={s.text_input}
                        underlineColorAndroid='transparent'
                        placeholder="Syötä pelin tunnus"
                        onChangeText={(game_id) => this.setState({game_id})}
                    />
                    <Button title='Liity' style={ s.button } onPress={() => { Alert.alert('Liity peliin', 'Liitytty peliin: ' + this.state.game_id) }}/>
                </View>
                <Button title='Luo uusi peli' style={ s.button } onPress={() => { this.createGame() }}/>
                <Button title='Näytä pelit' style={ s.button } onPress={() => { this.showGames() }}/>
            </View>
        );
    }
}

