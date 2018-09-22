import React from 'react';
import { StatusBar } from 'react-native'
import { createStackNavigator } from 'react-navigation';
import UserNameScreen from "./screens/UserNameScreen";
import SelectGameScreen from "./screens/SelectGameScreen";


const RootStack = createStackNavigator({
  UserNameScreen: UserNameScreen,
  SelectGameScreen: SelectGameScreen
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