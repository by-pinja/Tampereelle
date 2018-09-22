import React, { Component } from 'react';
import { Text, TextInput, View } from 'react-native';
import s from 'styles';

export default class UserNameScreen extends Component {
  constructor(props) {
    super(props);
    this.state = {text: ''};
  }

  render() {
    return (
      <View style={{padding: 10}}>
        <TextInput
          style={{height: 40}}
          placeholder="Anna käyttäjänimi"
          onChangeText={(text) => this.setState({text})}
        />
        <Button style={s.button}
            onPress={() => {
                Alert.alert('Nimi on ' + this.state.text);
            }}
            title="Ok"
        />
        <Text style={{padding: 10, fontSize: 42}}>
          {this.state.text.split(' ').map((word) => word && '🍕').join(' ')}
        </Text>
      </View>
    );
  }
}

