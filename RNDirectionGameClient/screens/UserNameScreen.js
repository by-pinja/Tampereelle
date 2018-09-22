import React, { Component } from 'react';
import { Text, TextInput, View, Button, Alert } from 'react-native';
import s from "../styles";

export default class UserNameScreen extends Component {
  constructor(props) {
    super(props);
    this.state = {text: ''};
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
                onChangeText={(text) => this.setState({text})}
            />
        </View>
        <Button style={s.button}
            onPress={() => {
                this.props.navigation.navigate("SelectGameScreen", {
                    player_name: this.state.text
                });
            }}
            color='#7439A2'
            title="Ok"
        />
      </View>
    );
  }
}

