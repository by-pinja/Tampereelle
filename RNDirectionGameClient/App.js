import React from 'react';
import { createStackNavigator } from 'react-navigation';
import {UserNameScreen} from 'UserNameScreen'


const RootStack = createStackNavigator({
  UserNameScreen: UserNameScreen
},
{
  initialRouteName: 'UserNameScreen',
});


export default class App extends React.Component {
  render() {
    return (
      <RootStack />
    );
  }
}