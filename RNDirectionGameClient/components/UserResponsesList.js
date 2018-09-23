import React, { Component } from 'react';
import {StyleSheet, View, Text, FlatList} from 'react-native';

export default class UserResponsesList extends Component {
    render() {
        const players = this.props.players;
        return (
            <View style={styles.container}>
                <FlatList
                    data={players}
                    keyExtractor={(item, index) => index.toString()}
                    renderItem={({item}) =><Text style={styles.item}>{item.player.name}: {Math.round(item.score * 100) / 100} Kulma: {Math.round(item.realAngle * 100) / 100}</Text>}
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
