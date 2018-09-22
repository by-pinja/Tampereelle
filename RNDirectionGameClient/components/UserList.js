import React, { Component } from 'react';
import {StyleSheet, View, FlatList} from 'react-native';

export default class UserList extends Component {
    render() {
        const players = this.props.players;
        return (
            <View style={styles.container}>
                <FlatList
                    data={players}
                    renderItem={({item}) => <Text style={styles.item} key={item.player_name}>{item.player_name}</Text>}
                />
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        paddingTop: 22
    },
    item: {
        padding: 6,
        fontSize: 16,
        height: 30,
    },
});
