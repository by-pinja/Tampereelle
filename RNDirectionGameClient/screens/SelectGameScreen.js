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
        fetch('https://techdays2018.appspot.com/api/games', {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            }
        }).then((response) => response.json()).then((responseJson) => {
            this.props.navigation.navigate("LobbyScreen", {
                game_id: responseJson.ID
            });
        });
    }
    showGame(id) {
        fetch('http://techdays2018.appspot.com/api/games/'+id, {
            method: 'GET',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json'
            }
        }).then((response) => response.json()).then((responseJson) => {
            Alert.alert("Message", JSON.stringify(responseJson))
        });
    }
    joinGame(id) {
        if(id){
            fetch('http://techdays2018.appspot.com/api/games/'+id, {
                method: 'POST',
                body: JSON.stringify({player_Name}),
                headers: {
                    Accept: 'application/json',
                    'Content-Type': 'application/json'
                }
            }).then((response) => response.json()).then((responseJson) => {
                Alert.alert("Message", JSON.stringify(responseJson))
            });
        }
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
                    <Button title='Liity' style={ s.button } onPress={() => { this.joinGame(this.state.game_id) }}/>
                </View>
                <Button title='Luo uusi peli' style={ s.button } onPress={() => { this.createGame() }}/>
                <View style={{paddingTop: 10}}>
                    <Button title='Näytä pelit' style={ s.button } onPress={() => { this.showGame(this.state.game_id) }}/>
                </View>
            </View>
        );
    }
}

