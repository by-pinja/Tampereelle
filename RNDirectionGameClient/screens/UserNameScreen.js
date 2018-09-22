import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import s from "../styles";
import { Constants, Permissions, Font } from 'expo';


export default class UserNameScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            player_name: ''
        };
    }
    render() {
        return (
          <View style={{padding: 10, backgroundColor: '#FFF', display: 'flex', flex: 1}}>
            <Text style={s.h1}>Tampereelle</Text>
            <View style={{paddingBottom: 10}}>
                <Text style={s.h2}>Pelaajan nimi:</Text>
                <TextInput
                    style={s.text_input}
                    underlineColorAndroid='transparent'
                    placeholder="Syötä pelaajan nimi"
                    onChangeText={(player_name) => this.setState({player_name})}
                />
            </View>
            <Button style={s.button}
                onPress={() => {
                    this.props.navigation.navigate("SelectGameScreen", {
                        player_name: this.state.player_name
                    });
                }}
                color='#7439A2'
                title="Ok"
            />
          </View>
        );
    }
}

